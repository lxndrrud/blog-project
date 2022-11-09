import React from 'react'
import { useDispatch, useSelector, useStore } from 'react-redux'
import MainLayout from '../layouts/MainLayout/MainLayout'
import PostList from '../components/Posts/PostsList/PostsList'
import { useEffect } from 'react'
import { fetchUserPosts } from '../store/actions/posts'
import { postsReducer } from '../store/reducers/posts'
import Swal from 'sweetalert2'

function UserPostsPage() {
    const dispatch = useDispatch()
    const store = useStore()
    const token = useSelector(state => state.users.token)

    useEffect(() => {
        dispatch(fetchUserPosts({ token }))
        .then(() => {
            const error = store.getState().posts.error
            if (error) {
                Swal.fire({
                    title: 'Произошла ошибка!',
                    text: error,
                    icon: "error"
                })
                dispatch(postsReducer.actions.clearError())
                return
            }
        })
    }, [])

    return (
        <MainLayout title={'Моё, родноё'} >
            <PostList />
        </MainLayout>
    )
}

export default UserPostsPage