import React from 'react'
import moment from 'moment'
import CustomLink from '../../UI/CustomLink/CustomLink'
import { useSelector } from 'react-redux'
import PostPictureImage from '../../UI/PostPictureImage/PostPictureImage'

function CheckPostItem({ post }) {
    const currentUser = useSelector(state => state.users.currentUser)
    return (
        <div className='mb-3 w-full p-2 flex flex-col
        rounded-md border border-solid border-[black]'>
            <h3 className='text-[28px] font-bold self-center'>{post.title}</h3>
            <h4 className='font-bold self-center'>Автор: {post.author_login}</h4>
            <PostPictureImage destination={`/detail/${post.id}`} altDescription={post.title} filepath={post.picture} />
            <h4 className='font-bold self-center'>Аннотация</h4>
            <p>{post.annotation}</p>
            <h4 className='font-bold self-center'>Содержимое</h4>
            <p>{post.text}</p>
            <div className='flex flex-col mt-2 text-[14px]'>
                <h4>Время создания: { moment(post.created_at).format('DD.MM.YYYY HH:mm').toString()} </h4>
                { post.time_start && <h4>Время начала доступности: {moment(post.time_start).format('DD.MM.YYYY HH:mm').toString()}</h4> }
                { post.time_end && <h4>Время конца доступности: {moment(post.time_end).format('DD.MM.YYYY HH:mm').toString()}</h4> }
            </div>
            <div className='flex mt-2 self-center'>
                <span className='mr-2'><CustomLink text={'Одобрить'} destination={`/approve/${post.id}`} /></span>
                <span className='mr-2'><CustomLink text={'Отклонить'} destination={`/reject/${post.id}`} /></span>
                { 
                    currentUser && currentUser.id === post.id_author
                    &&
                    <span><CustomLink text={'Удалить'} destination={`/deletePost/${post.id}`} /></span> 
                }
            </div>
        </div>
    )
}

export default CheckPostItem