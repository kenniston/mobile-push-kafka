package br.com.dynamiclight.push;

import java.util.Objects;

public class Push {
    private String message;

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    @Override
    public String toString() {
        return "Push(message='" + message + "')";
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Push push = (Push) o;
        return message.equals(push.message);
    }

    @Override
    public int hashCode() {
        return Objects.hash(message);
    }
}
