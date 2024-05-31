import { useNavigate } from "react-router-dom"
import { useMutation } from '@tanstack/react-query'
import useStore from '../store'
import { Credential, SignUp } from '../types'
import { useError } from './useError'
import axios from 'axios'

export const useMutateAuth = () => {
    const navigate = useNavigate()
    const resetEditedRival = useStore((state) => state.resetEditedRival)
    const { switchErrorHandling } = useError()

    // 第一引数に関数, 第二引数にレスポンスに合わせた処理記述
    const loginMutation = useMutation(
        async (user: Credential) => 
            await axios.post(`${process.env.REACT_APP_API_URL}/login`, user),
        {
            onSuccess: () => {
                navigate('/user')
            },
            onError: (err: any) => {
                // CSRFまわりのエラーだけはdata.messageにある
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message)
                } else {
                    switchErrorHandling(err.response.data)
                }
            },
        }
    )

    const registerMutation = useMutation(
        async (user: SignUp) => 
            await axios.post(`${process.env.REACT_APP_API_URL}/signup`, user),
        {
            onError: (err: any) => {
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message)
                } else {
                    switchErrorHandling(err.response.data)
                }
            },
        }
    )

    const logoutMutation = useMutation(
        async () => 
            await axios.post(`${process.env.REACT_APP_API_URL}/logout`),
        {
            onSuccess: () => {
                resetEditedRival()
                navigate('/')
            },
            onError: (err: any) => {
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message)
                } else {
                    switchErrorHandling(err.response.data)
                }
            },
        }
    )
    return { loginMutation, registerMutation, logoutMutation }
}