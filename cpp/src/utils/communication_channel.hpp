#ifndef COMMUNICATION_CHANNEL_H
#define COMMUNICATION_CHANNEL_H

#include <thread>
#include <mutex>
#include <condition_variable>
#include <list>

// Monitor much like Go channels. Inspired by https://st.xorian.net/blog/2012/08/go-style-channel-in-c/

template <class T>
class CommunicationChannel {
    public:
        void put(const T &e)
        {
            std::unique_lock<std::mutex> lock(m_mutex);
            m_queue.push_back(e);
            m_condition_variable.notify_one();
        }
        bool isEmpty()
        {
            std::unique_lock<std::mutex> lock(m_mutex);
            return m_queue.empty();
        }
        T get()
        {
            std::unique_lock<std::mutex> lock(m_mutex);
            m_condition_variable.wait(lock, [&](){ return !m_queue.empty(); });
            T result = m_queue.front();
            m_queue.pop_front();
            return result;
        }
    private:
        std::mutex m_mutex;
        std::condition_variable m_condition_variable;
        std::list<T> m_queue;
};

#endif // COMMUNICATION_CHANNEL_H