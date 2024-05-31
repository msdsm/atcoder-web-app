import { FC, memo } from 'react'
import { Table } from '../types'

const TableItemMemo: FC<Table> = ({
    id,
    atcoder_id,
    rating,
    streak,
}) => {
    return (
        <tr className="even:bg-gray-100">
            <td className="px-6 py-4">{atcoder_id}</td>
            <td className="px-6 py-4">{rating}</td>
            <td className="px-6 py-4">{streak}</td>
        </tr>
    )
}
// export const TableItem = memo(TableItemMemo)
export const TableItem = TableItemMemo