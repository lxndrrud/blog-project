import React from 'react'
import { useSelector } from 'react-redux'
import Preloader from '../../UI/Preloader/Preloader'
import CheckPostItem from '../CheckPostItem/CheckPostItem'
import PostItem from '../PostItem/PostItem'

function PostsList({ approvementFlag=false }) {
    let posts = useSelector(state => state.posts.posts)
    let isLoading = useSelector(state => state.posts.isLoading)
    return (
        <div className='w-full flex flex-col'>
            { 
                isLoading 
                ?
                    <Preloader />
                : (
                    approvementFlag
                    ?
                        posts && posts.length > 0 
                        ? 
                            posts.map(post => (
                                <CheckPostItem post={post} key={post.id} />
                            ))
                        : 
                            <span className='text-[20px] self-center'>Постов нет...</span>
                    :
                        posts && posts.length > 0
                        ?
                            posts.map(post => (
                                <PostItem post={post} key={post.id} />
                            ))
                        :
                            <span className='text-[20px] self-center'>Постов нет...</span>
                )
            }
        </div>
    )
}

export default PostsList