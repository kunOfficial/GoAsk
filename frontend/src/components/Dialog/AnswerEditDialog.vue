<template>
    <el-dialog v-model="dialogFormVisible" title="编辑回答" width="50%">
        <h2>{{ question?.title }}</h2>
        <p>{{ question?.description }}</p>
        <el-form :model="answerService" size="large">
            <el-form-item label="回答内容" label-width="80px">
                <el-input :rows="20" v-model="answerService.content" type="textarea" autocomplete="off" maxlength="1000"
                    show-word-limit />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button size="large" @click="$emit('cancel')">取消</el-button>
                <!-- TODO: 点击提交按钮后，回答post到后端 -->
                <el-button size="large" type="primary" @click="commitAnswer()">
                    提交
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import AnswerType from '@/types/Answer/AnswerType'
import AnswerService from '@/types/Answer/AnswerService';
import QuestionType from '@/types/Question/QuestionType'
import client from "@/utils/client"
import error from "@/utils/info"


onMounted(async () => {
    try {
        const response = await client.getQuestion(props.question_id)
        const body = response.data
        if (response.status === 200 && body.status === 200) {
            question.value = body.data
            // console.log(body.data)
        } else {
            error.ErrorAlert(body)
        }
    } catch (err) {
        console.log(err)
    }
})

const props = defineProps<{
    dialogFormVisible: boolean,
    // 当 priorAnswer 为 undefine 时，说明是第一次提交答案，而不是修改答案
    priorAnswer?: AnswerType
    question_id: number
}>()

const question = ref<QuestionType>()

const answerService = ref<AnswerService>({
    question_id: props.question_id,
    content: props.priorAnswer?.content
})

const emit = defineEmits<{
    (e: 'cancel'): void,
    (e: 'post-success', newAnswer: AnswerType): void
    (e: 'patch-success', newAnswer: AnswerType): void
}>()

// 如果是修改答案
if (props.priorAnswer != undefined) {
    answerService.value.content = props.priorAnswer.content
    answerService.value.answer_id = props.priorAnswer.id
}

// 提交答案
async function commitAnswer() {
    if (props.priorAnswer == undefined) {
        const response = await client.postAnswer(answerService.value);
        if (response.status == 200 && response.data.status == 200) {
            // TODO: 将 error 改为 info
            error.SuccessMessage("回答发布成功！")
            emit("post-success", response.data.data)
        } else {
            error.ErrorAlert(response.data)
            emit("cancel")
        }
    } else {
        client.patchAnswer(answerService.value)
        const response = await client.patchAnswer(answerService.value);
        if (response.status == 200 && response.data.status == 200) {
            error.SuccessMessage("回答修改成功！")
            emit("patch-success", response.data.data)
        } else {
            error.ErrorAlert(response.data)
            emit("cancel")
        }
    }
}


</script>

<style scoped>
p {
    margin-bottom: 50px;
}
</style>