import { createSlice } from "@reduxjs/toolkit";
import { fetchApprovedPosts } from "../actions/posts";


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
    }
})
