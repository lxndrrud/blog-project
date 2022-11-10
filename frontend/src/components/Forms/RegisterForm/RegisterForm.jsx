import React, { useState } from 'react'
import { useDispatch, useStore } from 'react-redux'
import { useNavigate } from 'react-router'
import { registerRequest } from '../../../store/actions/users'
import CustomButton from '../../UI/CustomButton/CustomButton'
import CustomInput from '../../UI/CustomInput/CustomInput'
import Swal from 'sweetalert2'
import { usersReducer } from '../../../store/reducers/users'

function RegisterForm() {
    const dispatch = useDispatch()
    const store = useStore()
    const navigate = useNavigate()

    let [login, setLogin] = useState(null)
    let [password, setPassword] = useState(null)

    const validate = () => {
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

    const sendRegisterRequest = (e) => {
        e.preventDefault()

        if (!validate()) return

        dispatch(registerRequest({ login, password }))
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
                    title: 'Регистрация выполнена успешно!',
                    icon: "success",
                    timer: 3000,
                })
                setTimeout(() => navigate('/login'), 3200)
            }
        })

    }

    return (
        <div className='flex flex-col justify-center'>
            <CustomInput type="text" text={login} callback={setLogin} description="Логин" />
            <CustomInput type="password" text={password} callback={setPassword} description="Пароль" />
            <CustomButton text={'Зарегистрироваться'} callback={sendRegisterRequest} />
        </div>
    )
}

export default RegisterForm