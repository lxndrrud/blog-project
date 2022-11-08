import React, { useEffect } from 'react'
import { useDispatch, useSelector, useStore } from 'react-redux'
import Swal from 'sweetalert2'
import PostsList from '../components/Posts/PostsList/PostsList'
import MainLayout from '../layouts/MainLayout/MainLayout'
import { fetchPostsNeedToApprove } from '../store/actions/posts'
import { postsReducer } from '../store/reducers/posts'

function ApprovementPostsPage() {
    const dispatch = useDispatch()
    const store = useStore()
    const token = useSelector(state => state.users.token)

    useEffect(() => {
        dispatch(fetchPostsNeedToApprove({ token }))
        .then(() => {
            const error = store.getState().posts.error
            if (error) {
                Swal.fire({
                    title: "Произошла ошибка!",
                    text: error,
                    icon: "error"
                })
            }
            dispatch(postsReducer.actions.clearError())
            return
        })
    }, [dispatch])

    return (
        <MainLayout title={'Посты для проверки'} >
            <PostsList approvementFlag />
        </MainLayout>
    )
}

export default ApprovementPostsPage