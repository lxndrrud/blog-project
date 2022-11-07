import React from 'react'
import { useDispatch, useStore } from 'react-redux'
import { useNavigate } from 'react-router'
import { useState } from 'react'
import Swal from 'sweetalert2'
import { loginRequest } from '../../../store/actions/users'
import CustomInput from '../../UI/CustomInput/CustomInput'
import CustomButton from '../../UI/CustomButton/CustomButton'
import { usersReducer } from '../../../store/reducers/users'

function LoginForm() {
    const dispatch = useDispatch()
    const store = useStore()
    const navigate = useNavigate()

    let [login, setLogin] = useState(null)
    let [password, setPassword] = useState(null)

    const validate = () => {
        console.log(login, password)
        if (!login) {
            Swal.fire({
                title: "Ошибка валидации",
                text: "Вы не ввели логин!",
                icon: "error",
            })
            return false
        }
        if (!password) {
            Swal.fire({
            title: "Ошибка валидации",
            text: "Вы не ввели логин!",
            icon: "error",
            })
            return false
        }
        return true
    } 

    const sendLoginRequest = (e) => {
        e.preventDefault()

        dispatch(loginRequest({ login, password }))
        .then(() => {
            const error = store.getState().users.error
            if (error) {
                Swal.fire({
                    title: 'Произошла ошибка!',
                    text: error,
                    icon: "error"
                })
                dispatch(usersReducer.actions.clearError())
                return
            } else {
                Swal.fire({
                    title: 'Вход выполнен успешно!',
                    icon: "success",
                    timer: 2000,
                })
                setTimeout(() => navigate('/'), 2200)
            }
        })

    }

    return (
        <div className='flex flex-col justify-center'>
            <CustomInput type="text" text={login} callback={setLogin} description="Логин" />
            <CustomInput type="password" text={password} callback={setPassword} description="Пароль" />
            <CustomButton text={'Вход'} callback={sendLoginRequest} />
        </div>
    )
}

export default LoginForm