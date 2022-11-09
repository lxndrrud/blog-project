import React from 'react'
import { useDispatch, useSelector, useStore } from 'react-redux'
import DialogForm from '../components/Forms/DialogForm/DialogForm'
import MainLayout from '../layouts/MainLayout/MainLayout'
import Swal from 'sweetalert2'
import { useNavigate, useParams } from 'react-router'
import { approvePostRequest } from '../store/actions/posts'
import { postsReducer } from '../store/reducers/posts'

function ApprovePostPage() {
    const dispatch = useDispatch()
    const store = useStore()
    const navigate = useNavigate()
    let { idPost } = useParams()
    idPost = parseInt(idPost)
    const token = useSelector(state => state.users.token)

    const approvePostCallback = (e) => {
        //e.preventDefault()

        dispatch(approvePostRequest({ token, idPost }))
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
                    title: 'Пост успешно одобрен!',
                    icon: 'success',
                    timer: 3000
                })
                setTimeout(navigate('/postsToCheck'), 3200)
            }
        })
    }

    return (
        <MainLayout title={'Одобрить публикацию поста'} >
            <DialogForm 
                question={'Пожалуйста, подтвердите, что Вы одобряете публикацию поста.'}  
                submitCallback={approvePostCallback}
                declineDestination={'/postsToCheck'}
            />
        </MainLayout>
    )
}

export default ApprovePostPage