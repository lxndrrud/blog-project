import { BrowserRouter, Route, Routes } from 'react-router-dom'
import IndexPage from './pages/IndexPage'
import LoginPage from './pages/LoginPage'
import './globals.css'
import PermissionMiddleware from './middlewares/PermissionMiddleware'
import CreatePostPage from './pages/CreatePostPage'
import ApprovementPostsPage from './pages/ApprovementPostsPage'

function App() {
    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={ <IndexPage /> } />
                    <Route path="/login" element={ <LoginPage /> } />
                    <Route path="/newPost" element={ 
                        <PermissionMiddleware permission={'СОЗД_ПОСТЫ'}>
                            <CreatePostPage />
                        </PermissionMiddleware> 
                    } />
                    <Route path="/postsToCheck" element={ 
                        <PermissionMiddleware permission={'МОДЕР_ПОСТЫ'}>
                            <ApprovementPostsPage />
                        </PermissionMiddleware> 
                    } />
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App
