import React from 'react'
import MainLayout from '../layouts/MainLayout/MainLayout'
import DialogForm from '../components/Forms/DialogForm/DialogForm'
import { useNavigate, useParams } from 'react-router'
import { useDispatch, useSelector, useStore } from 'react-redux'
import { deletePostRequest } from '../store/actions/posts'
import { postsReducer } from '../store/reducers/posts'
import Swal from 'sweetalert2'

function DeletePostPage() {
    const dispatch = useDispatch()
    const store = useStore()
    const navigate = useNavigate()

    const token = useSelector(state => state.users.token)
    let { idPost } = useParams()
    idPost = parseInt(idPost)
    const deletePostCallback = (e) => {
        //e.preventDefault()

        dispatch(deletePostRequest({ token, idPost }))
        .then(() => {
            const error = store.getState().posts.error
            if (error) {
                Swal.fire({
                    title: 'Произошла ошибка!',
                    text: error,
                    icon: 'error'
                })
                dispatch(postsReducer.actions.clearError())
                return
            } else {
                Swal.fire({
                    title: 'Пост успешно удалён!',
                    icon: 'success',
                    timer: 3000,
                })
                setTimeout(navigate('/'), 3200)
            }
        })
    }
    return (
        <MainLayout title={'Удаление поста'}>
            <DialogForm
                question={`Подтвердите удаление поста #${idPost}`}
                submitCallback={deletePostCallback}
                declineDestination={'/'}
            />
        </MainLayout>
    )
}

export default DeletePostPage