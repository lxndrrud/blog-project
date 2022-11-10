import React, { useEffect } from 'react'
import { useStore } from 'react-redux'
import { useNavigate } from 'react-router'
import RegisterForm from '../components/Forms/RegisterForm/RegisterForm'
import MainLayout from '../layouts/MainLayout/MainLayout'
import { loginCheck } from '../utils/LoginCheck'

function RegisterPage() {
    const store = useStore()
    const navigate = useNavigate()
    useEffect(() => {
        if (loginCheck(store)) {
            navigate('/')
            return
        }
    })

    return (
        <MainLayout title={'Регистрация пользователя'}>
            <RegisterForm />
        </MainLayout>
    )
}

export default RegisterPage