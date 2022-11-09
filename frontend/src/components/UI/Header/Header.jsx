import React from 'react'
import HeaderLink from '../HeaderLink/HeaderLink'
import { useDispatch, useStore } from 'react-redux'
import {loginCheck} from '../../../utils/LoginCheck'
import CustomButton from '../CustomButton/CustomButton'
import { usersReducer } from '../../../store/reducers/users'
import { useNavigate } from 'react-router'
import { permissionCheck } from '../../../utils/PermissionCheck'

function Header() {
    const dispatch = useDispatch()
    const store = useStore()
    const navigate = useNavigate()

    const logout = (e) => {
        e.preventDefault()

        dispatch(usersReducer.actions.logout())
        navigate("/")
    }

    return (
        <div className='pl-5 flex justify-start bg-[#e63946] text-[white]'>
            <div className='flex justify-between items-center'>
                <span className='p-1 mr-2 text-[30px]'>The Блог</span>
                <HeaderLink destination={'/'} text={'Посты'} />
                {
                    loginCheck(store)
                    ?
                        permissionCheck(store, 'СОЗД_ПОСТЫ')
                        ?
                            permissionCheck(store, 'МОДЕР_ПОСТЫ')
                            ?
                                [
                                    <HeaderLink destination={'/newPost'} text={'Написать'} />,
                                    <HeaderLink destination={'/myPosts'} text={'Мои посты'} />,
                                    <HeaderLink destination={'/postsToCheck'} text={'Модерировать'} />,
                                    <CustomButton text={"Выйти"} callback={logout}  />
                                ]
                            :
                                [
                                    <HeaderLink destination={'/newPost'} text={'Написать'} />,
                                    <HeaderLink destination={'/myPosts'} text={'Мои посты'} />,
                                    <CustomButton text={"Выйти"} callback={logout}  />
                                ]
                        :
                            [
                                <HeaderLink destination={'/myPosts'} text={'Мои посты'} />,
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