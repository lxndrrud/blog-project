import React from 'react'
import { useSelector } from 'react-redux'
import CheckPostItem from '../CheckPostItem/CheckPostItem'
import PostItem from '../PostItem/PostItem'

function PostsList({ approvementFlag=false }) {
    let posts = useSelector(state => state.posts.posts)
    return (
        <div className='w-full flex flex-col'>
            { 
                approvementFlag
                ?
                    posts && posts.map(post => (
                        <CheckPostItem post={post} key={post.id} />
                    ))
                :
                    posts && posts.map(post => (
                        <PostItem post={post} key={post.id} />
                    ))
            }
        </div>
    )
}

export default PostsList