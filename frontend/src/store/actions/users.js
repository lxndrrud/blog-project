import { createAsyncThunk } from "@reduxjs/toolkit"
import axios from "axios"

export const loginRequest = createAsyncThunk(
    "users/login",
    async ({ login, password }, thunkApi) => {
        try {
            console.log("action", login, password)
            const response = await axios.post("/backend/users/login", { login, password })
            return response.data
        } catch (error) {
            throw new Error(error?.response?.data.message || "Произошла непредвиденная ошибка!")
        }
    }
)
