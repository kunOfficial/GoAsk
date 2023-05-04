interface AnswerType {
    id: number,
    question_id: number,
    question_title: string,
    answerer_id: number,
    answerer_nick_name: string,
    content: string,
    updated_at: string,
    is_liked: boolean,
}

export default AnswerType;
