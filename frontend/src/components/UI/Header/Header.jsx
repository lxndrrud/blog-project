import React from 'react'
import HeaderLink from '../HeaderLink/HeaderLink'
import { useDispatch, useStore } from 'react-redux'
import {loginCheck} from '../../../utils/LoginCheck'
import CustomButton from '../CustomButton/CustomButton'
import { usersReducer } from '../../../store/reducers/users'
import { useNavigate } from 'react-router'

function Header() {
    const dispatch = useDispatch()
    const store = useStore()
    const navigate = useNavigate()

    const logout = (e) => {
        e.preventDefault()

        dispatch(usersReducer.actions.logout())
        .then(navigate("/"))
    }

    return (
        <div className='flex justify-start bg-[lightgreen]'>
            <div className='flex justify-between'>
                <HeaderLink destination={'/'} text={'Посты'} />
                {
                    loginCheck(store)
                    ?
                        [
                            <CustomButton text={"Выйти"} callback={logout}  />
                        ]
                    :
                        [
                            <HeaderLink destination={'/login'} text={'Вход'} />
                        ]
                }
            </div>
        </div>
    )
}

export default Header