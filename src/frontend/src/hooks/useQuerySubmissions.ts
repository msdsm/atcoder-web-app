import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { Submission } from '../types'
import { useError } from '../hooks/useError'

export const useQuerySubmissions = () => {
    const { switchErrorHandling} = useError()
    const getSubmissions = async () => {
        const { data } = await axios.get<Submission[]>(
            `${process.env.REACT_APP_API_URL}/user/submission`,
            { withCredentials: true}
        )
        return data
    }

    console.log("useQuerySubmissions")

    return useQuery<Submission[], Error>({
        queryKey: ['submissions'], // キャッシュのキー
        queryFn: getSubmissions,
        staleTime: Infinity,
        onError: (err : any) => {
            if(err.response.data.message) {
                switchErrorHandling(err.response.data.message)
            } else {
                switchErrorHandling(err.response.data)
            }
        },
    })
}