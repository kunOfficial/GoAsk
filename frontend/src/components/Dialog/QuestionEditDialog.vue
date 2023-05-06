<template>
    <el-dialog v-model="dialogFormVisible" title="编辑问题" >
        <el-form :model="questionService" size="large">
            <el-form-item label="问题标题" label-width="80px">
                <el-input v-model="questionService.title" />
            </el-form-item>
            <el-form-item label="问题描述" label-width="80px">
                <el-input v-model="questionService.description" type="textarea" autocomplete="off" maxlength="200"
                    show-word-limit />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button size="large" @click="$emit('cancel')">取消</el-button>
                <el-button size="large" type="primary" @click="postQuestion()">
                    提交
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import QuestionType from '@/types/Question/QuestionType'
import QuestionService from '@/types/Question/QuestionService'
import client from "@/utils/client"
import error from "@/utils/info"

const props = defineProps<{
    dialogFormVisible: boolean,
    priorQuestion?: QuestionType,
}>()

const emit = defineEmits<{
    // 这里一定要是 cancel，不然会有问题
    (e: 'cancel'): void,
    (e: 'success', newQuestion: QuestionType): void
}>()


const questionService = ref<QuestionService>({
    question_id: 0,
    title: "",
    description: "",
})
// 如果是修改问题
if (props.priorQuestion != undefined) {
    questionService.value.question_id = props.priorQuestion.id
    questionService.value.title = props.priorQuestion.title
    questionService.value.description = props.priorQuestion.description
}


async function postQuestion() {
    if (props.priorQuestion == undefined) {
        try {
            const response = await client.postQuestion(questionService.value);
            if (response.status == 200 && response.data.status == 200) {
                // TODO: 将 error 改为 info
                error.SuccessMessage("问题发布成功！")
                emit("success", response.data.data)
            } else {
                error.ErrorAlert(response.data)
                emit("cancel")
            }
        } catch (err) {
            console.log()
        }
    } else {
        const response = await client.patchQuestion(questionService.value);
        if (response.status == 200 && response.data.status == 200) {
            error.SuccessMessage("回答修改成功！")
            emit("success", response.data.data)
        } else {
            error.ErrorAlert(response.data)
            emit("cancel")
        }
    }
}
</script>