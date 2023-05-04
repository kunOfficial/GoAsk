import axios, { AxiosInstance, AxiosError, AxiosPromise, AxiosResponse } from "axios";
import UserLoginService from "@/types/User/UserSignService"
import UserType from "@/types/User/UserType"
import ResponseBodyType from "@/types/ResponseBodyType"
import store from '@/utils/store'
import AnswerType from "@/types/Answer/AnswerType";
import QuestionType from "@/types/Question/QuestionType";
import AnswerService from "@/types/Answer/AnswerService";
import QuestionService from "@/types/Question/QuestionService";

// https://github.com/axios/axios#config-defaults

const apiClient: AxiosInstance = axios.create({
    baseURL: "http://127.0.0.1:3000/api/v1",
    headers: {
        // "Origin":"localhost:5173",
        get: {
            "Content-type": "application/x-www-form-urlencoded",
        },
        patch: {
            "Content-type": "multipart/form-data",
            'Authorization': store.getToken(),
        }
    },
    validateStatus: function (status) {
        return status < 500; // 处理状态码小于500的情况
    }
});


// ################# User 用户操作 #################

function login(userService: UserLoginService): AxiosPromise<any> {
    return apiClient
        .post('/user/login', {
            "user_name": userService.user_name,
            "password": userService.password,
        })
}

function register(userService: UserLoginService): AxiosPromise<any> {
    return apiClient
        .post('/user/register', {
            "user_name": userService.user_name,
            "password": userService.password,
            "nick_name": userService.nick_name
        })
}

function getUserInfo(id: number): AxiosPromise<ResponseBodyType<UserType>> {
    return apiClient.get(`/users/${id}`)
}

function updateUserInfo(user: UserType): AxiosPromise<any> {
    // console.log(store.getToken())
    return apiClient
        .patch(`/users/${user.user_id}`, {
            "nick_name": user.nick_name,
            "about_me": user.about_me,
            "profession": user.profession,
        })
}

// 上传用户头像
function uploadAvatar(fd: FormData): AxiosPromise<ResponseBodyType<any>> {
    return apiClient.post(`/avatars`, fd, {
        headers: {
            'Authorization': store.getToken(),
        }
    })
}


// ############   Question 问题操作  #################

function getQuestions(page_size: number, page_num: number): AxiosPromise<ResponseBodyType<Array<QuestionType>>> {
    return apiClient.get(`/questions?page_num=${page_num}&page_size=${page_size}`)
}

function getUserQuestions(page_size: number, page_num: number, user_id: number): AxiosPromise<ResponseBodyType<Array<QuestionType>>> {
    return apiClient.get(`/questions?page_num=${page_num}&page_size=${page_size}&uid=${user_id}`)
}

function getQuestionAndAnswers(questionID: number): AxiosPromise<ResponseBodyType<QuestionType>> {
    return apiClient.get(`/questions/${questionID}?brief=false`)
}

function getQuestion(questionID: number): AxiosPromise<ResponseBodyType<QuestionType>> {
    return apiClient.get(`/questions/${questionID}?brief=true`)
}

function searchQuestions(filter: string, page_size: number, page_num: number): AxiosPromise<ResponseBodyType<Array<QuestionType>>> {
    return apiClient.get(`/questions/search?filter=${filter}&page_num=${page_num}&page_size=${page_size}`)
}



// postQuestion 发布问题
function postQuestion(questionService: QuestionService): AxiosPromise<ResponseBodyType<any>> {
    return apiClient.post(`/questions`, {
        "title": questionService.title,
        "description": questionService.description
    }, {
        headers: {
            // "Content-type": "multipart/form-data",
            'Authorization': store.getToken(),
        }
    })
}

// deleteQuestion 删除问题
function deleteQuestion(question_id: number): AxiosPromise<ResponseBodyType<any>> {
    return apiClient.delete(`questions/${question_id}`, {
        headers: {
            'Authorization': store.getToken(),
        }
    })
}
// patchQuestion 修改问题
function patchQuestion(questionService: QuestionService): AxiosPromise<ResponseBodyType<any>> {
    return apiClient.patch(`/questions/${questionService.question_id}`, {
        "title": questionService.title,
        "description": questionService.description
    }, {
        headers: {
            'Authorization': store.getToken(),
        }
    })
}



// ############   Anwer 回答操作  #################

// postAnswer 发布回答
function postAnswer(answerService: AnswerService): AxiosPromise<ResponseBodyType<AnswerType>> {
    return apiClient.post(`/answers`, {
        "content": answerService.content,
        "question_id": answerService.question_id
    }, {
        headers: {
            'Content-type': "application/json",
            'Authorization': store.getToken(),
        }
    })
}

// patchAnswer 修改回答
function patchAnswer(answerService: AnswerService): AxiosPromise<ResponseBodyType<any>> {
    return apiClient.patch(`/answers/${answerService.answer_id}`, {
        "content": answerService.content,
    }, {
        headers: {
            'Authorization': store.getToken(),
        }
    })
}


// getAnswers 获取一页回答
function getAnswers(page_size: number, page_num: number): AxiosPromise<ResponseBodyType<Array<AnswerType>>> {
    return apiClient.get(`/answers?page_num=${page_num}&page_size=${page_size}`)
}

// getAnswerLiked 获取是否点过赞
function getAnswerLikes(answer_id: number): AxiosPromise<ResponseBodyType<any>> {
    return apiClient.get(`/answers/${answer_id}/likes?uid=${store.getUserID()}`,
        {
            headers: {
                'Authorization': store.getToken(),
            }
        }
    )
}

// likeAnswer 点赞
function likeAnswer(answer_id: number): AxiosPromise<ResponseBodyType<any>> {
    if (store.isLogin()) {
        return apiClient.post(`/answers/${answer_id}/likes?uid=${store.getUserID()}`,
            {},
            {
                headers: {
                    'Authorization': store.getToken(),
                }
            })
    } else {
        return apiClient.post(`/answers/${answer_id}/likes`,
            {},
            {
                headers: {
                    'Authorization': store.getToken(),
                }
            })
    }
}

// cancelLikeAnswer 取消点赞
function cancelLikeAnswer(answer_id: number): AxiosPromise<ResponseBodyType<any>> {
    return apiClient.delete(`/answers/${answer_id}/likes?uid=${store.getUserID()}`)
}

// getUserAnswers 获取用户的回答
function getUserAnswers(page_size: number, page_num: number, user_id: number): AxiosPromise<ResponseBodyType<Array<AnswerType>>> {
    return apiClient.get(`/answers?page_num=${page_num}&page_size=${page_size}&uid=${user_id}`)
}

// deleteAnswer 删除回答
function deleteAnswer(aid: number): AxiosPromise<ResponseBodyType<any>> {
    return apiClient.delete(`/answers/${aid}`, {
        headers: {
            'Authorization': store.getToken(),
        }
    })

}


export default {
    register, login, getUserInfo, updateUserInfo,
    getAnswers, getQuestions, getQuestionAndAnswers,
    postAnswer, patchAnswer, postQuestion, deleteAnswer,
    patchQuestion, getUserQuestions, deleteQuestion, getQuestion, getUserAnswers, uploadAvatar,
    searchQuestions, likeAnswer, cancelLikeAnswer, getAnswerLikes
};