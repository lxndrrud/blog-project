import { createSlice } from "@reduxjs/toolkit";
import { approvePostRequest, createPostRequest, fetchApprovedPosts, fetchPostNeedToApprove, fetchPostsNeedToApprove } from "../actions/posts";


const initialState = {
    post: null,
    posts: [],
    isLoading: false,
    error: null,
}

const pending = (state) => {
    state.isLoading = true
}

const rejected = (state, action) =>  {
    state.isLoading = false
    if (action) state.error = action.error.message
}

const defaultFullfilled = (state) => {
    rejected(state, null)
}

export const postsReducer = createSlice({
    name: "posts", 
    initialState, 
    reducers: {
        clearError: (state, action) => {
            state.error = initialState.error
        }
    },
    extraReducers: {
        [fetchApprovedPosts.fulfilled]: (state, action) => {
            state.posts = action.payload.posts
            defaultFullfilled(state)
        },
        [fetchApprovedPosts.pending]: pending,
        [fetchApprovedPosts.rejected]: rejected,

        [createPostRequest.fulfilled]: (state, action) => {
            defaultFullfilled(state)
        },
        [createPostRequest.pending]: pending,
        [createPostRequest.rejected]: rejected,

        [fetchPostsNeedToApprove.fulfilled]: (state, action) => {
            state.posts = action.payload.posts
            defaultFullfilled(state)
        },
        [fetchPostsNeedToApprove.pending]: pending,
        [fetchPostsNeedToApprove.rejected]: rejected,

        [fetchPostNeedToApprove.fulfilled]: (state, action) => {
            state.post = action.payload.post
            defaultFullfilled(state)
        },
        [fetchPostNeedToApprove.pending]: pending,
        [fetchPostNeedToApprove.rejected]: rejected,

        [approvePostRequest.fulfilled]: (state, action) => {
            defaultFullfilled(state)
        },
        [approvePostRequest.pending]: pending,
        [approvePostRequest.rejected]: rejected,
    }
})
