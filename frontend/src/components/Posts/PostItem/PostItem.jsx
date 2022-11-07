import React from 'react'

function PostItem({ post }) {
    return (
        <div className='mb-3 w-full p-2 flex flex-col 
        rounded-md border border-solid border-[black]'>
            <h3>{post.title}</h3>
            <h4>{post.annotation}</h4>
            <div className='w-full h-[200px] justify-center align-center items-center'>Тут будет картинка</div>
            <h5>Просмотров: {post.views}</h5>
            <h5>Автор: {post.author_login}</h5>
        </div>
    )
}

export default PostItem