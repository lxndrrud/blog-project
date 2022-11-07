import { configureStore, combineReducers } from "@reduxjs/toolkit";
import { postsReducer } from "./reducers/posts"
import { usersReducer } from "./reducers/users";

const rootReducer = combineReducers({
    posts: postsReducer.reducer,
    users: usersReducer.reducer,
})

export const setupStore = () => configureStore({
    reducer: rootReducer,
})

