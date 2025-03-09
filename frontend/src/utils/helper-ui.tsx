import React from "react";

const mentionPattern = /([@#])\[([^\]]+)\]\(([^)]+)\)/g;

export const parseMentions = (text: string, handleClick: (type: string, id: string) => void): JSX.Element[] => {
    text = text.replaceAll("\n", "<br />")
    const parts = text.split(mentionPattern);
  
    // Process the parts to create React elements
    const elements: JSX.Element[] = [];
    let i = 0;
  
    while (i < parts.length) {
      const before = parts[i];
      const prefix = parts[i + 1]; // @ or #
      const name = parts[i + 2];
      const id = parts[i + 3];
  
      // Add text before mention
      if (before)
      elements.push(<span key={`text-${i}`}>{before.split("<br />").map((t, i) => <React.Fragment key={i}>{t}</React.Fragment>)}</span>);
  
      // Add mention element based on prefix
      if (prefix && name && id) {
        if (prefix === '@') {
          // Handle @mention for members
          elements.push(
            <a
              key={`${id}-${name}`}
              href="javascript:void(0);"
              onClick={() => handleClick("member", id)}
              className="mention-member"
            >
              @{name}
            </a>
          );
        } else if (prefix === '!') {
          // Handle #mention for clients
          elements.push(
            <a
              key={`${id}-${name}`}
              href="javascript:void(0);"
              onClick={() => handleClick("client", id)}
              className="mention-client"
            >
              #{name}
            </a>
          );
        } else if (prefix === '#') {
          // Handle #mention for clients
          elements.push(
            <a
              key={`${id}-${name}`}
              href="javascript:void(0);"
              onClick={() => handleClick("channel", id)}
              className="mention-client"
            >
              #{name}
            </a>
          );
        }
      }
  
      i += 4; // Move to the next group of parts
    }
  
    return elements;
  };
  
  