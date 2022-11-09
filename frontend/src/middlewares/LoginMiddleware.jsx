import React from 'react'
import { useSelector } from 'react-redux'
import { Navigate } from 'react-router-dom'

function LoginMiddleware({ children }) {
    const token = useSelector(state => state.users.token)
    if (!token) {
        return (
            <Navigate to={'/'} />
        )
    }
    return children
}

export default LoginMiddleware