import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { Rival } from '../types'
import useStore from '../store'
import { useError } from './useError'

export const useMutateRival = () => {
    const queryClient = useQueryClient()
    const { switchErrorHandling } = useError()
    const resetEditedRival = useStore((state) => state.resetEditedRival)

    // ライバル追加
    const createRivalMutation = useMutation(
        async (Rival: Omit<Rival, 'id'>) => {
            console.log("createRivalMutation")
            await axios.post<Rival>(`${process.env.REACT_APP_API_URL}/user/rival`, Rival)
        },
        {
            onSuccess: (res) => {
                const previousRivals = queryClient.getQueryData<Rival[]>(['rivals']) // キャッシュ
                if (previousRivals) {
                   // ここでRivalListコンポーネントの再レンダリングを走らせたい
                   queryClient.invalidateQueries(['rivals']); // キャッシュを無効化して再取得をトリガー
                   resetEditedRival();
                }
                resetEditedRival()
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

    // ライバル削除
    const deleteRivalMutation = useMutation(
        async (id: string) => {
            console.log("deleteRivalMutation")
            await axios.delete(`${process.env.REACT_APP_API_URL}/user/rival/${id}`)
        },
        {
            onSuccess: () => {
                const previousRivals = queryClient.getQueryData<Rival[]>(['rivals'])
                if (previousRivals) {
                    queryClient.invalidateQueries(['rivals'])
                }
                resetEditedRival()
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

    // 自分のatcoder idの変更
    const updateRivalMutation = useMutation(
        async (atcoder_id: string) => {
            console.log("updateRivalMutation")
            await axios.post(`${process.env.REACT_APP_API_URL}/user/profile/${atcoder_id}`)
        },
        {
            onSuccess: () => {
                const previousRivals = queryClient.getQueriesData<Rival[]>(['rivals'])
                if(previousRivals) {
                    queryClient.invalidateQueries(['rivals'])
                }
                resetEditedRival()
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
        createRivalMutation,
        deleteRivalMutation,
        updateRivalMutation,
    }
}