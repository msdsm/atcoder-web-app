import axios from 'axios'
import { useQuery } from '@tanstack/react-query'
import { Table } from '../types'
import { useError } from '../hooks/useError'

export const useQueryTables = () => {
    const { switchErrorHandling} = useError()
    const getTables = async () => {
        const { data } = await axios.get<Table[]>(
            `${process.env.REACT_APP_API_URL}/user/table`,
            { withCredentials: true}
        )
        return data
    }

    console.log("useQueryTables")

    return useQuery<Table[], Error>({
        queryKey: ['tables'], // キャッシュのキー
        queryFn: getTables,
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