import React from 'react'
import { useStore } from 'react-redux'
import { Navigate } from 'react-router-dom'
import { permissionCheck } from '../utils/PermissionCheck'

function PermissionMiddleware({ children, permissionCode }) {
    const store = useStore()
    if (!permissionCheck(store, permissionCode)) {
        return (
            <Navigate to={'/'} />
        )
    }
    return children
}

export default PermissionMiddleware