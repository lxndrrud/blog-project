import React from 'react'
import { useSelector } from 'react-redux'
import CustomLink from '../../UI/CustomLink/CustomLink'
import PostPictureLink from '../../UI/PostPictureLink/PostPictureLink'
import moment from 'moment'

function PostItem({ post }) {
    const currentUser = useSelector(state => state.users.currentUser)
    return (
        <div className='mb-3 w-full p-2 flex flex-col
        rounded-md shadow-xl border border-solid border-[#1d3557]'>
            <h3 className='text-[28px] font-bold self-center'>{post.title}</h3>
            <h4 className='self-center'>Автор: {post.author_login }</h4>
            <PostPictureLink destination={`/detail/${post.id}`} altDescription={post.title} filepath={post.picture} />
            <p className='mt-2 text-[20px] self-center'>{post.annotation}</p>
            <div className='flex flex-col mt-2 text-[14px]'>
                <h4>Просмотров: {post.views}</h4>
                <h4>Время создания: { moment(post.created_at).format('DD.MM.YYYY HH:mm').toString()} </h4>
                { post.time_start.Valid && <h4>Время начала доступности: {moment(post.time_start.Time).format('DD.MM.YYYY HH:mm').toString()}</h4> }
                { post.time_end.Valid && <h4>Время конца доступности: {moment(post.time_end.Time).format('DD.MM.YYYY HH:mm').toString()}</h4> }
            </div>
            { 
                currentUser && currentUser.id === post.id_author
                &&
                <div className='flex flex-col mt-2 text-[14px]'>
                    <span >{ post.approved ? "Одобрен" : "Не одобрен" }</span>
                    <span >{ post.rejected ? "Отклонен" : "Не отклонен" }</span>
                </div>
            }
            <div className='flex mt-2 self-center'>
                { 
                    currentUser && currentUser.id === post.id_author
                    &&
                    <span><CustomLink text={'Удалить'} destination={`/deletePost/${post.id}`} /></span> 
                }
            </div>
        </div>        
    )
}

export default PostItem