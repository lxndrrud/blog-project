import React from 'react'
import { useEffect } from 'react'
import { useStore } from 'react-redux'
import { useNavigate } from 'react-router'
import LoginForm from '../components/Forms/LoginForm/LoginForm'
import MainLayout from '../layouts/MainLayout/MainLayout'
import { loginCheck } from '../utils/LoginCheck'

function LoginPage() {
    const store = useStore()
    const navigate = useNavigate()
    useEffect(() => {
        if (loginCheck(store)) {
            navigate('/')
            return
        }
    })

    return (
        <MainLayout title={'Вход'}>
            <LoginForm />
        </MainLayout>
    )
}

export default LoginPage