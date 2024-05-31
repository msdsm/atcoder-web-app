import {create} from 'zustand'

type EditedRival = {
    id: number
    atcoder_id: string
}

type State = {
    editedRival: EditedRival // 状態管理したい対象としてEditedRival型の変数宣言
    updateEditedRival: (payload: EditedRival) => void // updateEditedRivalという関数の型宣言(入力:EditedRival型, 返り値:なし)
    resetEditedRival: () => void // 関数型宣言
}

// stateと関数の具体的な処理を追加
const useStore = create<State>((set) => ({
    editedRival: {id: 0, atcoder_id: ''},
    updateEditedRival: (payload) =>
        set({
            editedRival: payload,
        }),
    resetEditedRival: () => set({ editedRival: {id: 0, atcoder_id: ''} }),
}))

export default useStore