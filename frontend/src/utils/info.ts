import { ElMessage, ElMessageBox } from 'element-plus'
import ResponseBodyType from '@/types/ResponseBodyType'

function ErrorAlert(body: ResponseBodyType<any>) {
    ElMessageBox.alert(
        "error code:" + body.status + "\n" + body.msg,
        "错误"
    )
}

function SuccessMessage(msg: string) {
    ElMessage({
        type: 'success',
        message: msg,
    })
}

// 确认时，执行回调函数deleteCallback
function DeleteConfirm(msg: string, deleteCallback: () => void) {
    ElMessageBox.confirm(
        msg,
        'Warning',
        {
            confirmButtonText: 'OK',
            cancelButtonText: 'Cancel',
            type: 'warning',
        }
    )
        .then(() => {
            deleteCallback()
            ElMessage({
                type: 'success',
                message: '删除成功',
            })
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: '取消删除',
            })
        })
}

export default { ErrorAlert, SuccessMessage, DeleteConfirm };

// const open = () => {
//     ElMessageBox.confirm(
//         'proxy will permanently delete the file. Continue?',
//         'Warning',
//         {
//             confirmButtonText: 'OK',
//             cancelButtonText: 'Cancel',
//             type: 'warning',
//         }
//     )
//         .then(() => {
//             ElMessage({
//                 type: 'success',
//                 message: 'Delete completed',
//             })
//         })
//         .catch(() => {
//             ElMessage({
//                 type: 'info',
//                 message: 'Delete canceled',
//             })
//         })
// }