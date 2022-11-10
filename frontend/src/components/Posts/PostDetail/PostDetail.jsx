import React from 'react'
import { useSelector } from 'react-redux'
import PostPictureImage from '../../UI/PostPictureImage/PostPictureImage'
import Preloader from '../../UI/Preloader/Preloader'
import CustomLink from '../../UI/CustomLink/CustomLink'
import moment from 'moment'

function PostDetail() {
    let currentUser = useSelector(state => state.users.currentUser)
    let isLoading = useSelector(state => state.posts.isLoading)
    let post = useSelector(state => state.posts.post)
    return (
        <div className='w-full'>
        {
            isLoading
            ?
                <Preloader />
            :
                post 
                ?
                    <div className='mb-3 w-full p-2 flex flex-col
                    rounded-md border border-solid border-[black]'>
                        <h3 className='text-[28px] font-bold self-center'>{post.title}</h3>
                        <h4 className='self-center'>Автор: {post.author_login}</h4>
                        <PostPictureImage filepath={post.picture.String} altDescription={post.title} />
                        <h4 className='mt-2'>{post.annotation}</h4>
                        <p className='mt-2'>{post.text}</p>
                        <div className='flex flex-col mt-2 text-[14px]'>
                            <h4>Время создания: { moment(post.created_at).format('DD.MM.YYYY hh:mm').toString()} </h4>
                            { post.time_start.Valid && <h4>Время начала доступности: {moment(post.time_start.Time).format('DD.MM.YYYY hh:mm').toString()}</h4> }
                            { post.time_end.Valid && <h4>Время конца доступности: {moment(post.time_end.Time).format('DD.MM.YYYY hh:mm').toString()}</h4> }
                        </div>
                        <div className='flex mt-2 self-center'>
                            { 
                                currentUser && currentUser.id === post.id_author
                                &&
                                <span><CustomLink text={'Удалить'} destination={`/deletePost/${post.id}`} /></span> 
                            }
                        </div>
                    </div>
                : <span className='self-center'>Пост не найден...</span>
        }
        </div>
    )
}

export default PostDetail