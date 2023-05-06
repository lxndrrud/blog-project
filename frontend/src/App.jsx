import { BrowserRouter, Route, Routes } from 'react-router-dom'
import IndexPage from './pages/IndexPage'
import LoginPage from './pages/LoginPage'
import './globals.css'
import PermissionMiddleware from './middlewares/PermissionMiddleware'
import CreatePostPage from './pages/CreatePostPage'
import ApprovementPostsPage from './pages/ApprovementPostsPage'
import ApprovePostPage from './pages/ApprovePostPage'
import RejectPostPage from './pages/RejectPostPage'
import LoginMiddleware from './middlewares/LoginMiddleware'
import DeletePostPage from './pages/DeletePostPage'
import UserPostsPage from './pages/UserPostsPage'
import DetailApprovedPostPage from './pages/DetailApprovedPostPage'
import RegisterPage from './pages/RegisterPage'

function App() {
    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={ <IndexPage /> } />
                    <Route path="/login" element={ <LoginPage /> } />
                    <Route path="/register" element={ <RegisterPage /> } />
                    <Route path="/newPost" element={ 
                        <PermissionMiddleware permissionCode={'СОЗД_ПОСТЫ'}>
                            <CreatePostPage />
                        </PermissionMiddleware> 
                    } />
                    <Route path="/postsToCheck" element={ 
                        <PermissionMiddleware permissionCode={'МОДЕР_ПОСТЫ'}>
                            <ApprovementPostsPage />
                        </PermissionMiddleware> 
                    } />
                    <Route path="/approve/:idPost" element={
                        <PermissionMiddleware permissionCode={'МОДЕР_ПОСТЫ'}>
                            <ApprovePostPage />
                        </PermissionMiddleware>
                    } />
                    <Route path="/reject/:idPost" element={
                        <PermissionMiddleware permissionCode={'МОДЕР_ПОСТЫ'}>
                            <RejectPostPage />
                        </PermissionMiddleware>
                    } />
                    <Route path="/deletePost/:idPost" element={
                        <LoginMiddleware>
                            <DeletePostPage />
                        </LoginMiddleware>
                    } />
                    <Route path="/myPosts" element={
                        <LoginMiddleware>
                            <UserPostsPage />
                        </LoginMiddleware>
                    } />
                    <Route path="/detail/:idPost" element={
                        <DetailApprovedPostPage />
                    } />
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App
