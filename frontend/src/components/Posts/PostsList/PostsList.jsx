import React from 'react'
import { useSelector } from 'react-redux'
import PostItem from '../PostItem/PostItem'

function PostsList() {
    let posts = useSelector(state => state.posts.posts)
    return (
        <div className='w-full flex flex-col'>
            { posts && posts.map(post => (
                <PostItem post={post} key={post.id} />
            ))}
        </div>
    )
}

export default PostsList