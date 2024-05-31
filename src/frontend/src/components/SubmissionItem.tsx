import { FC, memo } from 'react'
import { Submission } from '../types'

const SubmissionItemMemo: FC<Submission> = ({
    atcoder_id,
    time,
    problem,
    diff,
}) => {
    return (
        <tr className="even:bg-gray-100">
            <td className="px-6 py-4">{atcoder_id}</td>
            <td className="px-6 py-4">{time}</td>
            <td className="px-6 py-4">{problem}</td>
            <td className="px-6 py-4">{diff}</td>
        </tr>
    )
}
export const SubmissionItem = memo(SubmissionItemMemo)