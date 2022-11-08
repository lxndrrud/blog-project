export const permissionCheck = (store, permission) => {
    return store.getState().users.permissions.findIndex(perm => perm.title === permission) !== -1
}