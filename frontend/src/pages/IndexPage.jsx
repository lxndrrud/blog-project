import React from 'react'
import { useEffect} from 'react'
import { useDispatch, useStore } from 'react-redux'
import MainLayout from '../layouts/MainLayout/MainLayout'
import { fetchApprovedPosts } from '../store/actions/posts'
import PostsList from '../components/Posts/PostsList/PostsList'
import Swal from 'sweetalert2'
import { postsReducer } from '../store/reducers/posts'

function IndexPage() {
    const dispatch = useDispatch()
    const store = useStore()
    useEffect(() => {
        dispatch(fetchApprovedPosts())
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
        <MainLayout title={'Самое интересное'}>
            <PostsList />
        </MainLayout>
    )
}

export default IndexPage