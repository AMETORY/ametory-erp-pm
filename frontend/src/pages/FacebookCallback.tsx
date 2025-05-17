import type { FC } from 'react';

interface FacebookCallbackProps {}

const FacebookCallback: FC<FacebookCallbackProps> = ({}) => {
    const urlParams = new URLSearchParams(window.location.search);
    const code = urlParams.get('code');
        return (<div>
            <h1>{code}</h1>
        </div>);
}
export default FacebookCallback;

