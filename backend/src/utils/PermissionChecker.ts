import { TPermission } from "../types/Permission.type";

export interface IPermissionChecker {
    canCreatePosts(permissionList: TPermission[]): boolean
    canModeratePosts(permissionList: TPermission[]): boolean
}

export class PermissionChecker implements IPermissionChecker {
    public canCreatePosts(permissionList: TPermission[]) {
        for (let permission of permissionList) {
            if (permission.code === "СОЗД_ПОСТЫ") return true
        }
        return false
    }

    public canModeratePosts(permissionList: TPermission[]) {
        for (let permission of permissionList) {
            if (permission.code === "МОДЕР_ПОСТЫ") return true
        }
        return false
    }
}
