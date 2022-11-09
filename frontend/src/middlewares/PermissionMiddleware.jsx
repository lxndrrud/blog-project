import React from 'react'
import { useSelector } from 'react-redux'
import { Navigate } from 'react-router-dom'

function checkPermissions(permissions, permission) {
    return permissions.findIndex(perm => perm.title === permission) !== -1
}

function PermissionMiddleware({ children, permission }) {
    const permissions = useSelector(state => state.users.permissions)
    if (!checkPermissions(permissions, permission)) {
        return (
            <Navigate to={'/'} />
        )
    }
    return children
}

export default PermissionMiddleware