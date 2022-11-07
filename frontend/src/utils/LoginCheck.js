
export const loginCheck = (store) => {
    if (store.getState().users.token) return true
    return false 
}