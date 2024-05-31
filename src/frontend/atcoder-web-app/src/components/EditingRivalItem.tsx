import { FC, memo } from 'react'
import { Rival } from '../types'

export const EditingRivalItem: FC<Omit<Rival, 'id'>> = ({
    atcoder_id: atcoder_id
}
) => {
    return (
        <tr className="even:bg-gray-100">
            <td className="px-6 py-4">{atcoder_id}</td>
        </tr>
    )
}