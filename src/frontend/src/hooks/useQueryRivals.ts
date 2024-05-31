import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { Rival } from '../types'
import { useError } from '../hooks/useError'

export const useQueryRivals = () => {
    const { switchErrorHandling} = useError()
    const getRivals = async () => {
        const { data } = await axios.get<Rival[]>(
            `${process.env.REACT_APP_API_URL}/user/rival`,
            { withCredentials: true}
        )
        return data
    }

    console.log("useQueryRivals")

    return useQuery<Rival[], Error>({
        queryKey: ['rivals'], // キャッシュのキー
        queryFn: getRivals,
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