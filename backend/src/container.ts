import { DatabaseConnection } from "./databaseConnection";
import { PermissionChecker } from "./utils/PermissionChecker";
import { RedisConnection } from "./redisConnection";
import { PermissionPostgresRepo } from "./repositories/Permission.repo";
import { PostPostgresRepo } from "./repositories/Post.repo";
import { UserPostgresRepo } from "./repositories/User.repo";
import { UserSessionRedisRepo } from "./repositories/UserSession.repo";
import { FileProcessor } from "./utils/FileProcessor";
import { Generator } from "./utils/Generator";
import { HttpErrorCreator } from "./utils/HttpErrorCreator";
import { UserService } from "./services/User.service";
import { PostService } from "./services/Post.service";
import { PermissionInfra } from "./infrastructure/Permission.infra";
import { UserController } from "./controllers/User.controller";
import { PostController } from "./controllers/Post.controller";

const generator = new Generator()
const permissionChecker = new PermissionChecker
const fileProcessor = new FileProcessor(generator)
const httpErrorCreator = new HttpErrorCreator()

const permissionRepo = new PermissionPostgresRepo(DatabaseConnection)
const postRepo = new PostPostgresRepo(DatabaseConnection)
const userRepo = new UserPostgresRepo(DatabaseConnection)
const userSessionRepo = new UserSessionRedisRepo(RedisConnection, generator)

const permissionInfra = new PermissionInfra(userSessionRepo, permissionRepo)

const userService = new UserService(
    userRepo, userSessionRepo, postRepo, permissionRepo, generator
)
const postService = new PostService(
    permissionInfra, permissionChecker, postRepo, userSessionRepo, fileProcessor
)

/*
async function test() {
    console.log(
        await generator.hashPassword('123')
    );
}
test()
*/

export const userController = new UserController(
    userService, httpErrorCreator
)
export const postController = new PostController(
    postService, httpErrorCreator
)
