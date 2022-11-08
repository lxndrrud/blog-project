import React from 'react'
import { useState } from 'react'
import { useDispatch, useSelector, useStore } from 'react-redux'
import { useNavigate } from 'react-router'
import CustomDatepicker from '../../UI/CustomDatepicker/CustomDatepicker'
import CustomInput from '../../UI/CustomInput/CustomInput'
import CustomButton from '../../UI/CustomButton/CustomButton'
import Swal from 'sweetalert2'
import { createPostRequest } from '../../../store/actions/posts'
import { postsReducer } from '../../../store/reducers/posts'
import CustomTextArea from '../../UI/CustomTextArea/CustomTextArea'

function CreatePostForm() {
    const dispatch = useDispatch()
    const store = useStore()
    const navigate = useNavigate()
    const token = useSelector(state => state.users.token)

    let [title, setTitle] = useState(null)
    let [text, setText] = useState(null)
    let [annotation, setAnnotation] = useState(null)
    let [timeStart, setTimeStart] = useState(null)
    let [timeEnd, setTimeEnd] = useState(null)

    const validate = () => {
        if (!title) {
            Swal.fire({
                title: 'Ошибка валидации!',
                text: 'Заголовок поста не указан!',
                icon: 'error'
            })
            return false
        }
        if (!annotation) {
            Swal.fire({
                title: 'Ошибка валидации!',
                text: 'Аннотация поста не указана!',
                icon: 'error'
            })
            return false
        }
        if (!text) {
            Swal.fire({
                title: 'Ошибка валидации!',
                text: 'Текст поста не указан!',
                icon: 'error'
            })
            return false
        }
        return true
    }

    const sendCreatePostRequest = (e) => {
        e.preventDefault()

        if (!(validate())) return

        dispatch(createPostRequest({ token, title, annotation, text, timeStart, timeEnd }))
        .then(() => {
            const error = store.getState().posts.error
            if (error) {
                Swal.fire({
                    title: "Произошла ошибка!",
                    text: error,
                    icon: "error"
                })
                dispatch(postsReducer.actions.clearError())
            } else {
                Swal.fire({
                    title: "Пост сохранен успешно!",
                    text: "Дождитесь подтверждения или отклонения от модераторов.",
                    icon: 'success',
                    timer: 3000
                })
                setTimeout(navigate('/'), 3200)
            }
        })
    }

    return (
        <div className='flex flex-col justify-center'>
            <CustomInput type="text" text={title} callback={setTitle} description="Заголовок" />
            <CustomInput type="text" text={annotation} callback={setAnnotation} description="Аннотация" />
            <CustomTextArea text={text} callback={setText} description="Основной текст" />
            <CustomDatepicker description={'Дата начала отображения'} callback={setTimeStart} />
            <CustomDatepicker description={'Дата конца отображения'} callback={setTimeEnd} /> 
            <CustomButton text={'Создать пост'} callback={sendCreatePostRequest} />
        </div>
    )
}

export default CreatePostForm