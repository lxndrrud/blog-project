import React from 'react'
import { useEffect } from 'react'
import { useDispatch,  useStore } from 'react-redux'
import { useParams } from 'react-router'
import MainLayout from '../layouts/MainLayout/MainLayout'
import { fetchApprovedPost } from '../store/actions/posts'
import { postsReducer } from '../store/reducers/posts'
import Swal from 'sweetalert2'
import PostDetail from '../components/Posts/PostDetail/PostDetail'

function DetailApprovedPostPage() {
    const dispatch = useDispatch()
    const store = useStore()
    let { idPost } = useParams()
    idPost = parseInt(idPost)

    useEffect(() => {
        dispatch(fetchApprovedPost({ idPost }))
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
        <MainLayout title="Подробная информация о посте">
            <PostDetail />
        </MainLayout>
    )
}

export default DetailApprovedPostPage