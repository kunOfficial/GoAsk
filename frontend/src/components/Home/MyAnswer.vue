<template>
    <div class="MyAnswer">
        <div class="body" @click="$router.push({ name: 'questions', params: { question_id: answer.question_id } })">
            <h2 class="questionTitle">{{ answer.question_title }}</h2>
            <p><span class="questioner">{{ answer.answerer_nick_name }}:</span>
                {{ answer.content }}
            </p>
            <p class="UpdatedAt">{{ answer.updated_at }}</p>
        </div>
        <EditButton class="EditAnswer" @click.stop="dialogFormVisible = true">修改回答</EditButton>
        <DeleteButton @click.stop="deleteAnswer">删除回答</DeleteButton>
        <AnswerEditDialog :priorAnswer="answer" :question_id="answer.question_id" :dialogFormVisible="dialogFormVisible"
            @cancel="(dialogFormVisible = false)"
            @patch-success="(a) => { dialogFormVisible = false; $emit('update-answer', a) }"
            @close="(dialogFormVisible = false)" />
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import AnswerType from '@/types/Answer/AnswerType';
import EditButton from '@/components/Common/EditButton.vue';
import AnswerEditDialog from '@/components/Dialog/AnswerEditDialog.vue';
import client from '@/utils/client'
import info from '@/utils/info'
import DeleteButton from '../common/DeleteButton.vue';

const props = defineProps<{
    answer: AnswerType,
}>()
const emit = defineEmits<{
    (e: 'delete-answer', aid: number): void
    (e: 'update-answer', answer: AnswerType): void
}>()

const dialogFormVisible = ref(false)

function deleteAnswer() {
    info.DeleteConfirm(`是否要删除问题"${props.answer.question_title}"的回答:\n"${props.answer.content.substring(0, 30)}"?`, async () => {
        const response = await client.deleteAnswer(props.answer.id)
        if (response.status == 200 && response.data.status == 200) {
            emit('delete-answer', props.answer.id)
        } else {
            info.ErrorAlert(response.data)
        }
    })
}
</script>


<style scoped>
.MyAnswer {
    background: rgb(220, 230, 243);
    /* border: 10px; */
    width: 100%;
    border-radius: 5px;
    position: relative;
}

.body {
    cursor: pointer;
}


.MyAnswer:hover {
    background: rgb(206, 223, 244);
}

h2 {
    margin-left: 10px;
}

.body .questioner {
    margin-left: 10px;
    font-size: 25px;
    color: grey;
    font-weight: bolder;
}

.body p {
    padding-right: 5px;
    padding-left: 5px;
}

.UpdatedAt {
    position: absolute;
    right: 10px;
    bottom: 0px;
    color: grey;
}

.EditAnswer {
    margin: 10px;
}
</style>
