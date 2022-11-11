import React from 'react'
import { useDispatch, useSelector, useStore } from 'react-redux'
import DialogForm from '../components/Forms/DialogForm/DialogForm'
import MainLayout from '../layouts/MainLayout/MainLayout'
import Swal from 'sweetalert2'
import { useNavigate, useParams } from 'react-router'
import { rejectPostRequest } from '../store/actions/posts'
import { postsReducer } from '../store/reducers/posts'

function RejectPostPage() {
    const dispatch = useDispatch()
    const store = useStore()
    const navigate = useNavigate()
    let { idPost } = useParams()
    idPost = parseInt(idPost)
    const token = useSelector(state => state.users.token)

    const rejectPostCallback = (e) => {
        //e.preventDefault()

        dispatch(rejectPostRequest({ token, idPost }))
        .then(() => {
            const error = store.getState().posts.error
            if (error) {
                Swal.fire({
                    title: 'Произошла ошибка',
                    text: error,
                    icon: error
                })
                dispatch(postsReducer.actions.clearError())
                return
            } else {
                Swal.fire({
                    title: 'Пост успешно отклонен!',
                    icon: 'success',
                    timer: 3000
                })
                setTimeout(navigate('/postsToCheck'), 3200)
            }
        })
    }

    return (
        <MainLayout title={'Отклонить публикацию поста'} >
            <DialogForm 
                question={'Пожалуйста, подтвердите, что Вы отклоняете публикацию поста.'}  
                submitCallback={rejectPostCallback}
                declineDestination={'/postsToCheck'}
            />
        </MainLayout>
    )
}

export default RejectPostPage