import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { useError } from '../hooks/useError'

export const useQueryProfile = () => {
    const { switchErrorHandling} = useError()
    const getProfile = async () => {
        const { data } = await axios.get<string>(
            `${process.env.REACT_APP_API_URL}/user/profile`,
            { withCredentials: true}
        )
        return data
    }
    return useQuery<string, Error>({
        queryKey: ['profile'], // キャッシュのキー
        queryFn: getProfile,
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