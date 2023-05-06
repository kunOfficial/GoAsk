<template>
    <div class="MyQuestion">
        <div class="body" @click="$router.push({ name: 'questions', params: { question_id: question.id } })">
            <h2>{{ question.title }}</h2>
            <p>{{ question.description }}</p>
            <p class="createdAt">{{ question.updated_at }}</p>
        </div>
        <div class="buttonSection">
            <EditButton class="EditQuestion" @click.stop="dialogFormVisible = true">修改问题</EditButton>
            <DeleteButton class="DeleteQuestion" @click.stop="deleteQuestion">删除问题</DeleteButton>
        </div>
        <QuestionEditDialog :dialogFormVisible="dialogFormVisible" :priorQuestion="question"
            @close="(dialogFormVisible = false)" @cancel="(dialogFormVisible = false)"
            @success="(q) => { dialogFormVisible = false; emit('update-question', q) }" />
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import QuestionType from '@/types/Question/QuestionType';
import EditButton from '@/components/Common/EditButton.vue';
import QuestionEditDialog from '@/components/Dialog/QuestionEditDialog.vue';
import DeleteButton from '../common/DeleteButton.vue';
import info from "@/utils/info"
import client from "@/utils/client"

const props = defineProps<{
    question: QuestionType
}>()


const emit = defineEmits<{
    (e: 'delete-question', qid: number): void
    (e: 'update-question', question: QuestionType): void
}>()


const dialogFormVisible = ref(false)

function deleteQuestion() {
    info.DeleteConfirm(`是否要删除问题"${props.question.title}"?`, async () => {
        const response = await client.deleteQuestion(props.question.id)
        if (response.status == 200 && response.data.status == 200) {
            emit('delete-question', props.question.id)
        } else {
            info.ErrorAlert(response.data)
        }
    })
}

</script>

<style scoped>
.MyQuestion {
    background-color: antiquewhite;
    border-radius: 10px;
    position: relative;
}

.body {
    cursor: pointer;
}

h2 {
    margin-left: 10px;
}

p {
    padding: 5px;
}

.createdAt {
    position: absolute;
    left: 10px;
    bottom: 0px;
    color: grey;
}

.buttonSection {
    text-align: right;
    padding: 10px;
}

.EditQuestion {
    margin-right: 10px;
}
</style>