import { createSlice } from "@reduxjs/toolkit";
import { loginRequest, registerRequest } from '../actions/users'

const initialState = {
    token: null,
    user: null,
    currentUser: null,
    permissions: [],
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

export const usersReducer = createSlice({
    name: "users", 
    initialState, 
    reducers: {
        clearError: (state, action) => {
            state.error = initialState.error
        },
        logout: (state, action) => {
            state.currentUser = initialState.currentUser
            state.token = initialState.token
        }
    },
    extraReducers: {
        [loginRequest.fulfilled]: (state, action) => {
            state.token = action.payload.token
            state.currentUser = action.payload.currentUser
            state.permissions = action.payload.permissions
            defaultFullfilled(state)
        },
        [loginRequest.pending]: pending,
        [loginRequest.rejected]: rejected,

        [registerRequest.fulfilled]: (state, action) => {
            defaultFullfilled(state)
        },
        [registerRequest.pending]: pending,
        [registerRequest.rejected]: rejected,
    }
})
