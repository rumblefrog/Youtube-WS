<template>
	<div id="Home" class="columns">
		<div class="column is-4 is-offset-4">
			<b-field>
				<b-input placeholder="Search Videos"
					type="search"
					expanded>
				</b-input>
				<p class="control">
					<button class="button is-primary">
						<font-awesome-icon icon="search" />
					</button>
				</p>
			</b-field>
		</div>
	</div>
</template>


<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src

@Component({
  components: {
    HelloWorld,
  },
})
export default class Home extends Vue {}
</script>
