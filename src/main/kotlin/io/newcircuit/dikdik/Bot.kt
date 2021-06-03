package io.newcircuit.dikdik

import io.newcircuit.dikdik.commands.*
import io.newcircuit.dikdik.config.Config
import io.newcircuit.dikdik.events.Interactions
import io.newcircuit.dikdik.events.Messages
import io.newcircuit.dikdik.models.ButtonState
import io.newcircuit.dikdik.models.ChannelMap
import io.newcircuit.dikdik.models.Command
import io.newcircuit.dikdik.models.Question
import org.javacord.api.DiscordApi
import org.javacord.api.DiscordApiBuilder
import org.javacord.api.interaction.ApplicationCommand
import org.javacord.api.interaction.ApplicationCommandBuilder
import java.util.concurrent.CompletableFuture


class Bot(val config: Config) {
    private val client: DiscordApi = DiscordApiBuilder()
        .setToken(config.token)
        .login().join()
    val commands = ArrayList<Command>()
    val channels = HashMap<Long, ChannelMap>()
    val clicks = ButtonState.getState()
    val votes = HashMap<Long, Question>()

    fun start() {
        registerEventListeners()
        registerCommands()
        println("Ready")
    }

    private fun registerEventListeners() {
        val msgEvent = Messages(this)
        val intEvent = Interactions(this)

        client.addListener(msgEvent)
        client.addListener(intEvent)
    }

    private fun registerCommands() {
        val testServer = client.getServerById(718433475828645928).get()
        val cmds = arrayListOf(
            Joke(this),
            Fact(this),
            TalkIn(this),
            Stop(this),
            Button(this),
            Ask(this),
            CloseVote(this),
        )

        val globalCmds = client.globalApplicationCommands.join()
        val hasCmd = fun(name: String): Boolean {
            for (cmd in globalCmds) {
                if (cmd.name == name) {
                    return true
                }
            }
            return false
        }

        for (command in cmds) {
            val builder = ApplicationCommandBuilder()
                .setName(command.name)
                .setDescription(command.description)
            val options = command.getOptions()
            for (option in options) {
                builder.addOption(option.build())
            }
            println("Registering: ${command.name}")
            if (!hasCmd(command.name)) {
                builder.createForServer(testServer).join()
            }
            builder.createGlobal(client).join()
            this.commands.add(command)
        }
    }
}
