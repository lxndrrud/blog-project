import { BrowserRouter, Route, Routes } from 'react-router-dom'
import IndexPage from './pages/IndexPage'
import LoginPage from './pages/LoginPage'
import './globals.css'

function App() {
    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={ <IndexPage /> } />
                    <Route path="/login" element={ <LoginPage /> } />
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App
