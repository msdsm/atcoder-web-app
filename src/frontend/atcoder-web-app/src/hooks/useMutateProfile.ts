import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { useError } from './useError'

export const useMutateProfile = () => {
    const queryClient = useQueryClient()
    const { switchErrorHandling } = useError()
    // 自分のatcoder idの変更
    const updateProfileMutation = useMutation(
        async (atcoder_id: string) => {
            console.log("updateProfileMutation")
            await axios.post(`${process.env.REACT_APP_API_URL}/user/profile/${atcoder_id}`)
        },
        {
            onSuccess: () => {
                const previousProfile = queryClient.getQueriesData<string>(['profile'])
                if(previousProfile) {
                    queryClient.invalidateQueries(['profile'])
                }
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

    return {
        updateProfileMutation,
    }
}