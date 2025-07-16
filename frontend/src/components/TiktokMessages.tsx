import type { FC } from 'react';

interface TiktokMessagesProps {
    sessionId: string;
}

const TiktokMessages: FC<TiktokMessagesProps> = ({
    sessionId
}) => {
        return (
            <h1>Hallo</h1>
        );
}
export default TiktokMessages;