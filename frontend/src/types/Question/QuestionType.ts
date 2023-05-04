import AnswerType from "../Answer/AnswerType";

interface QuestionType {
    id: number,
    questioner_id: number,
    title: string,
    description: string,
    view: number,
    updated_at: string,
    questioner_nick_name: string,
    answers: AnswerType[] 
}

export default QuestionType;
